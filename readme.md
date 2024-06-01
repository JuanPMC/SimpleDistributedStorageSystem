# Programming Assignment 2

## Overview

This program consists of two entities: the indexing server and the nodes. The nodes are unable to discover each other without the presence of an indexing server.

## Prerequisites

- The deployment scripts will automatically check and install Go.
- A Linux environment is required.

## Usage

1. **Deploy Indexing Server:**

   Navigate to the 'deployment' directory and run the following command:

   bash ./deployIndexing_server.sh

   This will start the indexing server. **DO NOT STOP THE INDEXING SERVER or the nodes won't be able to find each other.**

2. **Deploy Peer Nodes:**

   - For Node 1:
     bash ./deploy_node1.sh

   - For Node 2:
     bash ./deploy_node2.sh

   - For Node 3:
     bash ./deploy_node3.sh

   - For Node 4:
     bash ./deploy_node4.sh

   Run the appropriate deployment script for each peer node.

3. **Using the nodes:**
When you deploy a node, this terminal will become a client, it has a friendly menue to list the available files and download files. You have to write the full name of the path.


## Troubleshooting

- If everything does not work as expected, try re-deploying everything in order.
- Examine log files for more detailed information on any issues.
