# rpi-infrared-scanner
 Allows you to scan codes sent by infrared remotes.


### Usage
Before building the application, make sure to change the following code (In line 12)
```go
const pinNum uint8 = 4
```
The application uses the `BCM` layout of the pins.

### Build
Execute following command to cross-compile the code to a binary which is understandable by the *Raspberry-Pi*.

 ```
 make build
 ```
 After the application's binary has been built, transfer it to the *Raspberry-Pi* using a utility like `scp`.

 ### Usage
 After the binary has been transferred to the *Raspberry-Pi*, run it by using following command:
 ```
 ./scanner
 ```

 ### Going further
 Of course a scanner alone is not useful, instead the code can be modified to represent some sort of listener which can trigger events on the raspberry pi when a certain code is recognized.