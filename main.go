package main

import (
	"fmt"
	"os"

	"github.com/EmeraldLS/image-conversion/conversion"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "initial",
		Short: "File Conversion Program",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("All about file conversion. use -h to see list of commands")
		},
	}
	pngCommand := &cobra.Command{
		Use:   "png-to-jpeg",
		Short: "Program to convert png image to jpeg",
		Long:  "",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fileName := args[0]
			outputFileName := args[1]

			imgBytes, err := os.ReadFile(fmt.Sprintf("./src/%v.png", fileName))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			message, err := conversion.PngToJpeg(imgBytes, outputFileName)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(message)
		},
	}

	JpegCommand := &cobra.Command{
		Use:   "jpeg-to-png",
		Short: "Program to convert jpeg images to png",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fileName := args[0]
			outputFileName := args[1]

			imgBytes, err := os.ReadFile(fmt.Sprintf("./src/%v.jpg", fileName))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = conversion.ToPng(imgBytes, outputFileName)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("image converted successfully")
		},
	}

	rootCmd.AddCommand(pngCommand)
	rootCmd.AddCommand(JpegCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
