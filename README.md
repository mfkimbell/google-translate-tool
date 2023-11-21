# google-translate-tool

**Tools used:**
* `Go Routines` for concurrency of task execution
* `Channels` for communication of Go Routines
* `Wait Groups` sync mechanism that blocks code to allow for other functions to complete execution, similar to `await` keyword


One thing I think is important to understand for this project is that it utilizes concurrency through Go Routines, aka tasks are always being actively worked on when there is downtime during API calls and things alike, however, there is no paralellism, such as when you have dynamically scaling worker servers/containers. We use a channel to 

Here is the output with no flags:
<img width="606" alt="Screenshot 2023-11-21 at 5 13 00 PM" src="https://github.com/mfkimbell/google-translate-tool/assets/107063397/eb1c4960-0783-4e28-9d06-f1b09d820268">

Here is the output with correct flags:
<img width="376" alt="Screenshot 2023-11-21 at 5 15 17 PM" src="https://github.com/mfkimbell/google-translate-tool/assets/107063397/bc9fb26b-3ea6-4419-b606-d0ec533eee75">
