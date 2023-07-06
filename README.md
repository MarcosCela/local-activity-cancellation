# Reproduce


Execute the script, it will:

- Create a temporalite server.
- Create a worker that connects to that temporalite server.
- Start a workflow, which starts a local activity. The local activity simulates a failure (SIGINT).

Sometimes (race condition?) The workflow shows up as "canceled" instead of timedout. Need
more investigation to get a consistent cancelled error.


```shell
go run .                                                                                                                                                                                                                                                                           130 ↵
2023/07/06 18:17:36 INFO  No logger configured for temporal client. Created default one.
2023/07/06 18:17:36 INFO  Started Worker Namespace default TaskQueue worker WorkerID 51113@COMP-YV2D730KLG@
{"level":"warn","time":"2023-07-06T18:17:36+02:00","message":"Starting the workflow"}
2023/07/06 18:17:36 WARN  Starting the local activity, wish me luck Namespace default TaskQueue worker WorkerID 51113@COMP-YV2D730KLG@ WorkflowType MyWorkflow WorkflowID Test RunID aa545df2-81f6-478b-a4b1-cf1ac0ccb227 Attempt 1
{"level":"warn","time":"2023-07-06T18:17:36+02:00","message":"Handling signal: terminated"}
2023/07/06 18:17:37 INFO  Stopped Worker Namespace default TaskQueue worker WorkerID 51113@COMP-YV2D730KLG@
2023/07/06 18:17:37 ERROR The activity was cancelled Namespace default TaskQueue worker WorkerID 51113@COMP-YV2D730KLG@ ActivityID 1 ActivityType MyActivity Attempt 1 WorkflowType MyWorkflow WorkflowID Test RunID aa545df2-81f6-478b-a4b1-cf1ac0ccb227
panic: workflow execution error (type: , workflowID: Test, runID: aa545df2-81f6-478b-a4b1-cf1ac0ccb227): canceled

goroutine 1 [running]:
main.main()
	/Users/marcos.cela/GIT/MarcosCela/local-activity-cancellation/main.go:68 +0x3a8
exit status 2
```