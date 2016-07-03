package main

import (
	"runtime"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/pflag"
)

// App encapsulates all of the parameters necessary for starting up
// an aws mock metadata server. These can either be set via command line or directly.
type App struct {
	AvailabilityZone string
	AppPort          string
	PublicIp         string
	PublicHostname   string
	PrivateIp        string
	Hostname         string
	InstanceID       string
	InstanceType     string
	MacAddress       string
	VpcID            string
	SecurityGroups   string
	RoleArn          string
	RoleName         string
	Verbose          bool
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	app := &App{}
	app.addFlags(pflag.CommandLine)
	pflag.Parse()

	if app.Verbose {
		log.SetLevel(log.DebugLevel)
	}

	app.NewServer()
}

func (app *App) addFlags(fs *pflag.FlagSet) {
	fs.StringVar(&app.AvailabilityZone, "availability-zone", app.AvailabilityZone, "Availability zone")
	fs.StringVar(&app.AppPort, "app-port", app.AppPort, "Http port")
	fs.StringVar(&app.PublicIp, "public-ip", app.PublicIp, "Public IP")
	fs.StringVar(&app.PublicHostname, "public-hostname", app.PublicHostname, "Public Hostname")
	fs.StringVar(&app.PrivateIp, "private-ip", app.PrivateIp, "Private IP")
	fs.StringVar(&app.Hostname, "hostname", app.Hostname, "EC2 instance hostname")
	fs.StringVar(&app.InstanceID, "instance-id", app.InstanceID, "EC2 instance ID")
	fs.StringVar(&app.InstanceType, "instance-type", app.InstanceType, "EC2 instance type")
	fs.StringVar(&app.MacAddress, "mac-address", app.MacAddress, "MAC address")
	fs.StringVar(&app.VpcID, "vpc-id", app.VpcID, "VPC ID")
	fs.StringVar(&app.SecurityGroups, "security-groups", app.SecurityGroups, "Security groups")
	fs.StringVar(&app.RoleArn, "role-arn", app.RoleArn, "IAM role ARN")
	fs.StringVar(&app.RoleName, "role-name", app.RoleName, "IAM role name")
	fs.BoolVar(&app.Verbose, "verbose", false, "Verbose")
}
