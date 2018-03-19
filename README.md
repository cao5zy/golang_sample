# golang_sample

Samples for using go

[inherit value from other object](struct_learning.go)  

[reference items from other file but within the same package](lib1)  

Interface in go looks a bit different to other languages. For detailed usage, please refer to [about_interface](about_interface)  
The best practice to query interface in Go is as following.  

    if animal, ok := who.(Animal); ok { //query the Animal interface
		animal.Walk()
	}

To run the app, there are two ways:  

    go run *.go
or  

    go build *.go  
    ./[package file] //run the root package which has the same name with the folder

