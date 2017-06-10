enum Result {
  OK = 0,
  ERROR = 2,
}

struct Cmd {
  1: string cmdLine = "";
  2: i64 ticket = 0;
}

exception ExecuteException {
  1: string what;
  2: string output;
}

service Parallel {
   // round trip
   string Ping();

   // execute a shell command:
   string Execute(1:Cmd command) 
        throws (1:ExecuteException e);
}
