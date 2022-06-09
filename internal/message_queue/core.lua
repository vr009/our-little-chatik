local fiber = require('fiber')
local uuid = require('uuid')

box.cfg{
    listen = 3301,
}


box.schema.space.create('msg_queue',
        {
            if_not_exists = true,
            format = { { 'chat_id', type = 'uuid' },
                       { 'msg_id', type = 'uuid' , unique = true },
                       { 'sender_id', type = 'uuid' },
                       { 'receiver_id', type = 'uuid' },
                       { 'payload', type = 'string' },
                       { 'created_at', type = 'number' } }
        })

box.space.msg_queue:create_index('chat_index', {unique = true, if_not_exists = true, parts = { { 1, 'uuid' }, {6, 'number'} }})

box.space.msg_queue:create_index('msg_index', {if_not_exists = true, parts = { { 2, 'uuid' }}})

local queue = {}

-- in this table we keep chat_ids and receiver_ids with the value of last unread message
local chats = {}

-- in this table we keep chat updates for receiver_id
local chats_upd = {}

-- sync thing
queue._wait = fiber.channel()

function queue.put(chat_id, sender_id, receiver_id, payload)
    local msg_id = uuid()
    local created_at = os.time()

    print(chat_id, sender_id, receiver_id, payload)

    if chats[chat_id] == nil then
        chats[chat_id] = {}
    end

    if not chats[chat_id][receiver_id] or chats[chat_id][receiver_id] == -1 then
        chats[chat_id][receiver_id] = created_at
    end

    -- we put the id of the last message for chat list update
    if not chats_upd[chat_id] then
        chats_upd[chat_id] = {}
    end
    chats_upd[chat_id] = msg_id:str()

    return box.space.msg_queue:insert{
        uuid.fromstr(chat_id),
        msg_id,
        uuid.fromstr(sender_id),
        uuid.fromstr(receiver_id),
        payload, created_at}
end


function queue.take_new_messages_from_space(chat_id, since_msg_id, sender_id, receiver_id)
    local since = 0

    if not chats[chat_id] then
        return {}
    end

    since = chats[chat_id][receiver_id]
    chats[chat_id][receiver_id] = -1

    local batch = {}
    for _, tuple in box.space.msg_queue.index.chat_index:pairs({ uuid.fromstr(chat_id) }) do
        if (since ~=-1 and since <= tuple[6]) or since_msg_id == nil then
            table.insert(batch, tuple)
        end
    end
    return batch
end


function queue.fetch_chat_list_update(chat_list)
    local batch = {}
    for _, chat_id in ipairs(chat_list) do
        local msg_id = chats_upd[chat_id]
        if msg_id ~= nil then
            local tuple = box.space.msg_queue.index.msg_index:get(uuid.fromstr(msg_id))
            table.insert(batch, tuple)
        end
    end
    return batch
end



rawset(_G, 'put', queue.put)
rawset(_G, 'take_msgs', queue.take_new_messages_from_space)
rawset(_G, 'fetch_chats_upd', queue.fetch_chat_list_update)

box.once('debug', function() box.schema.user.grant('guest', 'super') end)

box.schema.user.create('test', {password = 'test'})
box.schema.user.grant('test', 'execute', 'universe')
box.schema.user.grant('test', 'read,write', 'space', 'msg_queue')

return queue

-- uuid.fromstr('2743f114-524d-4e3b-8ed0-20666e976d39')