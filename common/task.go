package common

type Task struct {
    Id int64
    Name string
    Status bool
    Command string
    Args string
    Runner string
    Version int
}

type HeartbeatReq struct {
    Host    string
    Tasks   map[int64]int  // key: task id, value: task version
}

type HeartbeatTask struct {
    Task  *Task
    Delete bool
}

type HeartbeatResp struct {
    Host  string
    Tasks map[int64]*HeartbeatTask
}
