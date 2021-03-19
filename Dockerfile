FROM tarantool/tarantool
COPY init.lua .
CMD ["tarantool", "init.lua"]
