package tests

//
//import (
//	"context"
//	"github.com/docker/docker/api/types/container"
//	"github.com/makeopensource/leviathan/docker"
//	"github.com/rs/zerolog"
//	"github.com/rs/zerolog/log"
//	"os"
//	"testing"
//)
//
//// TestHelloName calls greetings.Hello with a name, checking
//// for a valid return value.
//func TestHelloName(t *testing.T) {
//	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
//
//	web_gen, err := docker.NewLocalClient()
//	if err != nil {
//		log.Fatal().Msg("Failed to setup local docker web_gen")
//	}
//
//	//web_gen, err := docker.NewSSHClient("r334@192.168.50.123")
//	//if err != nil {
//	//	log.Fatal().Msg("Failed to setup docker web_gen")
//	//}
//
//	log.Info().Msg("Connected to remote web_gen")
//
//	err = docker.BuildImageFromDockerfile(web_gen, "../example/ex-Dockerfile", "testimage:latest")
//	if err != nil {
//		log.Error().Err(err).Msg("Failed to build image")
//		return
//	}
//
//	images, err := docker.ListImages(web_gen)
//	if err != nil {
//		log.Error().Msg("Failed to build image")
//		return
//	}
//
//	for _, image := range images {
//		log.Info().Msgf("Container names: %v", image.RepoTags)
//	}
//
//	newContainerId, err := docker.CreateNewContainer(
//		web_gen,
//		"92912992939",
//		"testimage:latest",
//		[]string{"py", "/home/autolab/student.py"},
//		container.Resources{
//			Memory:   512 * 1000000,
//			NanoCPUs: 2 * 1000000000,
//		},
//	)
//	if err != nil {
//		log.Error().Err(err).Msg("Failed to create container")
//		return
//	}
//
//	err = docker.CopyToContainer(web_gen, newContainerId, "../example/student/test.py")
//	if err != nil {
//		log.Error().Err(err).Msg("Failed to copy to container")
//	}
//
//	err = docker.StartContainer(web_gen, newContainerId)
//	if err != nil {
//		log.Error().Err(err).Msg("Failed to start container")
//		return
//	}
//
//	err, _ = docker.TailContainerLogs(context.Background(), web_gen, newContainerId)
//	if err != nil {
//		log.Error().Err(err).Msg("Failed to tail logs")
//		return
//	}
//
//	data, err := docker.ListContainers(web_gen)
//	if err != nil {
//		log.Error().Msg("Failed to build image")
//		return
//	}
//
//	for _, info := range data {
//		log.Info().Msgf("Container names: %v", info.Names)
//	}
//
//	//t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
//
//}
//
//// TestHelloEmpty calls greetings.Hello with an empty string,
//// checking for an error.
//func TestHelloEmpty(t *testing.T) {
//}
