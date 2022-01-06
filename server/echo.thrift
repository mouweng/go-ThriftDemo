namespace go echo

struct EchoReq {
    1: string msg;
}

struct EchoRes {
    1: string msg;
}

struct Num {
    1:required i32 id;
}

service Echo {
    EchoRes echo(1: EchoReq req);
    Num Add(1: Num num1, 2: Num num2);
}
