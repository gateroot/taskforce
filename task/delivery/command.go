package delivery

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"taskforce/task"
)

type AddOptions struct {
	Args struct {
		Title       string `required:"true"`
		Description string `required:"true"`
	} `positional-args:"yes" required:"yes"`
}

type StartOptions struct {
	Args struct {
		ID int `required:"true"`
	} `positional-args:"yes" required:"yes"`
}

type StopOptions struct {
	Args struct {
		ID int `required:"true"`
	} `positional-args:"yes" required:"yes"`
}

type PauseOptions struct {
	Args struct {
		ID int `required:"true"`
	} `positional-args:"yes" required:"yes"`
}

type CompleteOptions struct {
	Args struct {
		ID int `required:"true"`
	} `positional-args:"yes" required:"yes"`
}

type CloseOptions struct {
	Args struct {
		ID int `required:"true"`
	} `positional-args:"yes" required:"yes"`
}

type DeleteOptions struct {
	Args struct {
		ID int `required:"true"`
	} `positional-args:"yes" required:"yes"`
}

type GetOptions struct {
	Args struct {
		ID int `required:"true"`
	} `positional-args:"yes" required:"yes"`
}

type ListOptions struct{}

type Options struct {
	Add      AddOptions      `command:"add"  description:"add function"`
	Start    StartOptions    `command:"start" description:"start function"`
	Stop     StopOptions     `command:"stop" description:"stop function"`
	Pause    PauseOptions    `command:"pause" description:"pause function"`
	Complete CompleteOptions `command:"complete" description:"complete function"`
	Close    CloseOptions    `command:"close" description:"close function"`
	Delete   DeleteOptions   `command:"delete" description:"delete function"`
	Get      GetOptions      `command:"get" description:"get function"`
	List     ListOptions     `command:"list" description:"list function"`
}

type CommandLineDelivery struct {
	usecase task.Usecase
	view    task.ViewTaskUsecase
}

func NewCommandLineDelivery(usecase task.Usecase, view task.ViewTaskUsecase) *CommandLineDelivery {
	return &CommandLineDelivery{usecase, view}
}

func (cd *CommandLineDelivery) Exec() {
	opts := Options{}
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "taskforce"
	_, err := parser.Parse()
	if err != nil {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	switch parser.Active.Name {
	case "add":
		err := cd.usecase.New(opts.Add.Args.Title, opts.Add.Args.Description)
		if err != nil {
			fmt.Fprintf(os.Stdout, "add task failed.\n%v", err)
			os.Exit(1)
		}
	case "start":
		err := cd.usecase.Start(opts.Start.Args.ID)
		if err != nil {
			fmt.Fprintf(os.Stdout, "start task failed.\n%v", err)
			os.Exit(1)
		}
	case "stop":
		err := cd.usecase.Stop(opts.Stop.Args.ID)
		if err != nil {
			fmt.Fprintf(os.Stdout, "stop task failed.\n%v", err)
			os.Exit(1)
		}
	case "pause":
		err := cd.usecase.Pause(opts.Pause.Args.ID)
		if err != nil {
			fmt.Fprintf(os.Stdout, "pause task failed.\n%v", err)
			os.Exit(1)
		}
	case "complete":
		err := cd.usecase.Complete(opts.Complete.Args.ID)
		if err != nil {
			fmt.Fprintf(os.Stdout, "complete task failed.\n%v", err)
			os.Exit(1)
		}
	case "close":
		err := cd.usecase.Close(opts.Close.Args.ID)
		if err != nil {
			fmt.Fprintf(os.Stdout, "close task failed.\n%v", err)
			os.Exit(1)
		}
	case "delete":
		err := cd.usecase.Delete(opts.Delete.Args.ID)
		if err != nil {
			fmt.Fprintf(os.Stdout, "delete task failed.\n%v", err)
			os.Exit(1)
		}
	case "get":
		t, err := cd.view.Get(opts.Get.Args.ID)
		if err != nil {
			fmt.Fprintf(os.Stdout, "close task failed.\n%v", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "|%-5s|%-20s|%-50s|%-10s|\n", "ID", "TITLE", "DESCRIPTION", "STATE")
		fmt.Fprintf(os.Stdout, "|%-5d|%-20s|%-50s|%-10s|\n", t.GetId(), t.GetTitle(), t.GetDescription(), t.GetState().String())
	case "list":
		tasks, err := cd.view.List()
		if err != nil {
			fmt.Fprintf(os.Stdout, "close task failed.\n%v", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "|%-5s|%-20s|%-50s|%-10s|\n", "ID", "TITLE", "DESCRIPTION", "STATE")
		for _, t := range tasks {
			fmt.Fprintf(os.Stdout, "|%-5d|%-20s|%-50s|%-10s|\n", t.GetId(), t.GetTitle(), t.GetDescription(), t.GetState().String())
		}
	}
}
