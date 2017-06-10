enum Result {
  OK = 0,
  ERROR = 2,
}

struct Cmd {
  1: string cmdLine = "";
  2: i64 ticket = 0;
}

struct Output {
  1: string stdout;
  2: string stderr;
  3: map<string, string> tags;
}

exception ExecuteException {
  1: string what;
  2: Output output;
}

service Parallel {
   // round trip
   string Ping();

   // execute a shell command:
   Output Execute(1:Cmd command) 
            throws (1:ExecuteException e);
}
