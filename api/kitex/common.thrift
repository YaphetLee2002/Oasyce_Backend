exception Error {
    1: i32 HTTPCode,
    2: string Code,
    3: string Message,
    4: map<string, string> Data (go.tag = 'json:"Data,omitempty"'),
}