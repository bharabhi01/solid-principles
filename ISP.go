/*
	Interface Segregation Principle says that the client should not be forced to depend on interfaces they do not use.
*/

package main

import "fmt"

/*
	-----------------------Bad Example--------------------------------
	In the below example, we have a WorkerBad interface that has a Work, Eat and Sleep method.
	We have a HumanWorkerBad struct that implements the WorkerBad interface.
	We have a RobotWorkerBad struct that implements the WorkerBad interface.

	But, the RobotWorkerBad struct cannot implement the Eat and Sleep method.
	So, this violates the Interface Segregation Principle.

	If we RobotWorkerBad does not implement the Eat and Sleep method, then it will not be able to
	implement the WorkerBad interface.
*/

type WorkerBad interface {
	Work() string
	Eat() string
	Sleep() string
}

type HumanWorkerBad struct{}

func (hw HumanWorkerBad) Work() string {
	return "Human is working"
}

func (hw HumanWorkerBad) Eat() string {
	return "Human is eating"
}

func (hw HumanWorkerBad) Sleep() string {
	return "Human is sleeping"
}

type RobotWorkerBad struct{}

func (rw RobotWorkerBad) Work() string {
	return "Robot is working"
}

func (rw RobotWorkerBad) Eat() string {
	return "Robot is eating" // This is wrong because a robot cannot eat
}

func (rw RobotWorkerBad) Sleep() string {
	return "Robot is sleeping" // This is wrong because a robot cannot sleep
}

/*
	-----------------------Good Example--------------------------------
	In the below example, we have segregated the Worker interface into Worker, Eater and Sleeper interfaces.
	Due to this, now the RobotWorkerBad struct cannot implement the Eat and Sleep method.
	So, this satisfies the Interface Segregation Principle.
*/

type Worker interface {
	Work() string
}

type Eater interface {
	Eat() string
}

type Sleeper interface {
	Sleep() string
}

type HumanWorker struct{}

func (hw HumanWorker) Work() string {
	return "Human is working"
}

func (hw HumanWorker) Eat() string {
	return "Human is eating"
}

func (hw HumanWorker) Sleep() string {
	return "Human is sleeping"
}

type RobotWorker struct{}

func (rw RobotWorker) Work() string {
	return "Robot is working"
}

func ISP_Demo() {
	fmt.Println("Interface Segregation Principle")

	humanWorker := HumanWorker{}
	robotWorker := RobotWorker{}

	humanWorker.Work()
	humanWorker.Eat()
	humanWorker.Sleep()

	robotWorker.Work()
}
