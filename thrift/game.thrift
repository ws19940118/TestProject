namespace go  thrift


struct User {
    1: string name
    2: i64 age
}

struct Teacher {
    1: string name
    2: i64 age
}


struct gReq {
   1: User user
}

struct gResp {
    1: Teacher msg
}

service Game {
    gResp GetInfo(1: gReq req)
}