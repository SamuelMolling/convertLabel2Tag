# convertLabel2Tag
Code for converting MongoDB Atlas labels to tags in dedicated clusters.

It will create tags based on the cluster labels. If the label does not exist, it will ignore the error and move on to the next one.

## How to use
1. Export your public and private key

```bash
export ATLAS_PUBLIC_KEY=<your-public-key>
export ATLAS_PRIVATE_KEY=<your-private-key>
```

2. Run the script 

```bash
go run main.go
```

teste