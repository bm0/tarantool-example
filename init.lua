box.cfg({
	listen = 3301
})

local user = box.schema.create_space('user')
user:create_index('primary', {parts={{1, 'uint'}}})

function get_user(userID)
	return user:get(userID)
end
