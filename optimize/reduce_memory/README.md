https://chris124567.github.io/2021-06-21-go-performance/



```
optimize/reduce_memory] $go build -gcflags='-m -m'
# github.com/ludwig125/ludwig125_gosample/optimize/reduce_memory
./reduce.go:30:6: cannot inline foo: unhandled op DEFER
./reduce.go:84:6: cannot inline main: unhandled op DEFER
./reduce.go:85:27: inlining call to os.Create func(string) (*os.File, error) { var os..autotmp_3 *os.File; os..autotmp_3 = <N>; var os..autotmp_4 error; os..autotmp_4 = <N>; os..autotmp_3, os..autotmp_4 = os.OpenFile(os.name, int(578), os.FileMode(438)); return os..autotmp_3, os..autotmp_4 }
./reduce.go:100:14: inlining call to fmt.Println func(...interface {}) (int, error) { var fmt..autotmp_3 int; fmt..autotmp_3 = <N>; var fmt..autotmp_4 error; fmt..autotmp_4 = <N>; fmt..autotmp_3, fmt..autotmp_4 = fmt.Fprintln(io.Writer(os.Stdout), fmt.a...); return fmt..autotmp_3, fmt..autotmp_4 }
./reduce.go:98:14: inlining call to fmt.Println func(...interface {}) (int, error) { var fmt..autotmp_3 int; fmt..autotmp_3 = <N>; var fmt..autotmp_4 error; fmt..autotmp_4 = <N>; fmt..autotmp_3, fmt..autotmp_4 = fmt.Fprintln(io.Writer(os.Stdout), fmt.a...); return fmt..autotmp_3, fmt..autotmp_4 }
./reduce.go:103:13: inlining call to fmt.Println func(...interface {}) (int, error) { var fmt..autotmp_3 int; fmt..autotmp_3 = <N>; var fmt..autotmp_4 error; fmt..autotmp_4 = <N>; fmt..autotmp_3, fmt..autotmp_4 = fmt.Fprintln(io.Writer(os.Stdout), fmt.a...); return fmt..autotmp_3, fmt..autotmp_4 }
./reduce.go:103:55: cannot inline main.func1: function too complex: cost 118 exceeds budget 80
./reduce.go:17:7: can inline glob..func1 with cost 10 as: func() interface {} { b := make([]byte, 256); return &b }
./reduce.go:25:7: can inline glob..func2 with cost 74 as: func() interface {} { return sha256.New() }
./reduce.go:26:20: inlining call to sha256.New func() hash.Hash { var sha256.d·2 *sha256.digest; sha256.d·2 = <N>; sha256.d·2 = new(sha256.digest); sha256.d·2.Reset(); return hash.Hash(sha256.d·2) }
./reduce.go:69:11: make([]byte, 0, 256) does not escape
./reduce.go:103:28: int(testing.AllocsPerRun(100, func literal)) escapes to heap:
./reduce.go:103:28:   flow: ~arg1 = &{storage for int(testing.AllocsPerRun(100, func literal))}:
./reduce.go:103:28:     from int(testing.AllocsPerRun(100, func literal)) (spill) at ./reduce.go:103:28
./reduce.go:103:28:     from ~arg0, ~arg1 = <N> (assign-pair) at ./reduce.go:103:13
./reduce.go:103:28:   flow: {storage for []interface {} literal} = ~arg1:
./reduce.go:103:28:     from []interface {} literal (slice-literal-element) at ./reduce.go:103:13
./reduce.go:103:28:   flow: fmt.a = &{storage for []interface {} literal}:
./reduce.go:103:28:     from []interface {} literal (spill) at ./reduce.go:103:13
./reduce.go:103:28:     from fmt.a = []interface {} literal (assign) at ./reduce.go:103:13
./reduce.go:103:28:   flow: {heap} = *fmt.a:
./reduce.go:103:28:     from fmt.Fprintln(io.Writer(os.Stdout), fmt.a...) (call parameter) at ./reduce.go:103:13
./reduce.go:103:14: "Allocs:" escapes to heap:
./reduce.go:103:14:   flow: ~arg0 = &{storage for "Allocs:"}:
./reduce.go:103:14:     from "Allocs:" (spill) at ./reduce.go:103:14
./reduce.go:103:14:     from ~arg0, ~arg1 = <N> (assign-pair) at ./reduce.go:103:13
./reduce.go:103:14:   flow: {storage for []interface {} literal} = ~arg0:
./reduce.go:103:14:     from []interface {} literal (slice-literal-element) at ./reduce.go:103:13
./reduce.go:103:14:   flow: fmt.a = &{storage for []interface {} literal}:
./reduce.go:103:14:     from []interface {} literal (spill) at ./reduce.go:103:13
./reduce.go:103:14:     from fmt.a = []interface {} literal (assign) at ./reduce.go:103:13
./reduce.go:103:14:   flow: {heap} = *fmt.a:
./reduce.go:103:14:     from fmt.Fprintln(io.Writer(os.Stdout), fmt.a...) (call parameter) at ./reduce.go:103:13
./reduce.go:100:15: "Test FAIL" escapes to heap:
./reduce.go:100:15:   flow: ~arg0 = &{storage for "Test FAIL"}:
./reduce.go:100:15:     from "Test FAIL" (spill) at ./reduce.go:100:15
./reduce.go:100:15:     from ~arg0 = <N> (assign-pair) at ./reduce.go:100:14
./reduce.go:100:15:   flow: {storage for []interface {} literal} = ~arg0:
./reduce.go:100:15:     from []interface {} literal (slice-literal-element) at ./reduce.go:100:14
./reduce.go:100:15:   flow: fmt.a = &{storage for []interface {} literal}:
./reduce.go:100:15:     from []interface {} literal (spill) at ./reduce.go:100:14
./reduce.go:100:15:     from fmt.a = []interface {} literal (assign) at ./reduce.go:100:14
./reduce.go:100:15:   flow: {heap} = *fmt.a:
./reduce.go:100:15:     from fmt.Fprintln(io.Writer(os.Stdout), fmt.a...) (call parameter) at ./reduce.go:100:14
./reduce.go:98:15: "Test PASS" escapes to heap:
./reduce.go:98:15:   flow: ~arg0 = &{storage for "Test PASS"}:
./reduce.go:98:15:     from "Test PASS" (spill) at ./reduce.go:98:15
./reduce.go:98:15:     from ~arg0 = <N> (assign-pair) at ./reduce.go:98:14
./reduce.go:98:15:   flow: {storage for []interface {} literal} = ~arg0:
./reduce.go:98:15:     from []interface {} literal (slice-literal-element) at ./reduce.go:98:14
./reduce.go:98:15:   flow: fmt.a = &{storage for []interface {} literal}:
./reduce.go:98:15:     from []interface {} literal (spill) at ./reduce.go:98:14
./reduce.go:98:15:     from fmt.a = []interface {} literal (assign) at ./reduce.go:98:14
./reduce.go:98:15:   flow: {heap} = *fmt.a:
./reduce.go:98:15:     from fmt.Fprintln(io.Writer(os.Stdout), fmt.a...) (call parameter) at ./reduce.go:98:14
./reduce.go:98:15: "Test PASS" escapes to heap
./reduce.go:98:14: []interface {} literal does not escape
./reduce.go:100:15: "Test FAIL" escapes to heap
./reduce.go:100:14: []interface {} literal does not escape
./reduce.go:103:14: "Allocs:" escapes to heap
./reduce.go:103:28: int(testing.AllocsPerRun(100, func literal)) escapes to heap
./reduce.go:103:55: func literal does not escape
./reduce.go:103:13: []interface {} literal does not escape
./reduce.go:19:3: b escapes to heap:
./reduce.go:19:3:   flow: ~r0 = &b:
./reduce.go:19:3:     from &b (address-of) at ./reduce.go:20:10
./reduce.go:19:3:     from &b (interface-converted) at ./reduce.go:20:10
./reduce.go:19:3:     from return &b (return) at ./reduce.go:20:3
./reduce.go:19:12: make([]byte, 256) escapes to heap:
./reduce.go:19:12:   flow: b = &{storage for make([]byte, 256)}:
./reduce.go:19:12:     from make([]byte, 256) (spill) at ./reduce.go:19:12
./reduce.go:19:12:     from b := make([]byte, 256) (assign) at ./reduce.go:19:5
./reduce.go:19:3: moved to heap: b
./reduce.go:19:12: make([]byte, 256) escapes to heap
./reduce.go:26:20: new(sha256.digest) escapes to heap:
./reduce.go:26:20:   flow: sha256.d·2 = &{storage for new(sha256.digest)}:
./reduce.go:26:20:     from new(sha256.digest) (spill) at ./reduce.go:26:20
./reduce.go:26:20:     from sha256.d·2 = new(sha256.digest) (assign) at ./reduce.go:26:20
./reduce.go:26:20:   flow: ~R0 = sha256.d·2:
./reduce.go:26:20:     from hash.Hash(sha256.d·2) (interface-converted) at ./reduce.go:26:20
./reduce.go:26:20:     from ~R0 = <N> (assign-pair) at ./reduce.go:26:20
./reduce.go:26:20:   flow: ~r0 = ~R0:
./reduce.go:26:20:     from sha256.New() (interface-converted) at ./reduce.go:26:20
./reduce.go:26:20:     from return sha256.New() (return) at ./reduce.go:26:3
./reduce.go:26:20: new(sha256.digest) escapes to heap
<autogenerated>:1: .this does not escape
./reduce.go:74:36: index bounds check elided
```
