# Trusted Text

Trusted text aims to be a distributed ecosystem for creating, sharing and updating text on the web. The intent is to be one of the tools which helps combat the ongoing loss of trust in most existing forms of information. 

Content hosted using trusted text will be modifiable, But not secretly, and the original source will always be available.

## Key objectives

* It is possible for an author to claim a segment of trusted text, by signing it with an identity key (such as PGP). 
* It is possible to identify a segment of trusted text using a set of keys, identified by the author at creation
    - It is allowable to have a trusted text segment without any keys. This is considered as the 'profile' of the original author. It is suggested that outside amendments to is segment is ignored.
* It is possible to easily include this text in a webpage
* Multiple segments of text can refer to the same set of keys, but only one of them has the 'HEAD' flag
* In its simplest state, only the authenticated original author can move the 'HEAD' flag
* There exists a publically accessible repository of any set of published keys
* There is some form of distributed content delivery network, which will always return the consensus 'HEAD' for any given set of keys 

## Secondary objectives

* Mechanism for clear distinction between opinion segments (which the author always has full control over) and information segments, which can have a consensus control over the head pointer.
* Mechanism for a consensus to be formed on alternative interpretations of text. 
    - Possibly through using some kind of proof of stake. This should be linked to some idea of thorough research.... Not quite sure how to price that in though.
    - Or making use of some kind of outside metric of trustworthiness on certain fact checkers etc
    - (Danger) allowing users to set their consensus pillars, generating a dynamic view of the 'HEAD' Using the aggregate votes of only your pillars. 
    - Creating a momentum system, such that users who have their edits or original content sustained have more influence.       


## Sibling projects

Most rich text news content includes the use of images. While we could rely on the incongruence of the image with its caption, it would be preferable if we could also have a similar trusted principle with the images.

* This would need to be based on the original idea, where the author can post the original image and sign it. But this one should likely be immutable, and other resolutions or formats can be obtained on demand, and signed by the host of the original image. 
* Then the article can refer to the image with it's hash or unique key. This means that if the original writer wishes to change their image with an edit, they can do so. But this would then still be available for the past views of the page.


Using this method to validate videos is likely much more challenging, but it could be used quite easily to keep track of the signing and hashes of the files hosted elsewhere. 


## Deviation From 'traditional blockchain' 

Rather than store all the suggestions in an ordered list (prone to races) instead this is modelled as a 'cloud' of suggestions in an unordered map. In this environment, only the head hash matters. Each suggestion needs only keep track of the head hash at the time of it's creation. And branching is permitted. 

When a hash is promoted to a head hash, it is appended to the map of head hashes. All the blocks which are identified here must be retained with each replica of the chain, but the cloud of un-promoted blocks are optional. 

Thus we have a cloud of suggestions, out of which a definitive path is slowly cemented. This path can be a tree, as the consensus reverts to a previous node to continue. Analogous to an iterative algorithm backtracking to avoid local maxima/minima.


## Distribution strategy

The text chains will be distributed as a network of peers. One chain will exist as the founder. 

That chain can be called with a clone request, where it responds with the serialised chain. That new chain is then eligible as a peer, and will enter the peer list.

On a regular basis, peers will perform a handshake, identifying to each other if there are any missing hashes between them. If any are missing, they share with the peer which URL they can call to claim that resource. 

Peers also regularly check on the n and n+1 peers. By requesting a peerlist, you will be given the peerlist endpoints for all of their peers. After excluding yourself, a given node can check that that peer is a valid one, which is in alignment. 

You can identify specific 'pillars' which you trust, then your own node will only align itself with changes which are aligned to one or more of those nodes. In this manner you can assemble a web of trust which will allow valid messages to propagate through the network without allowing the same access to compromised ones. 

### Distributing complete blocks

The planned usage strategy would be that one would generate new blocks in a local environment. This will then be the only environment that your private key needs to be present, and trusted text should integrate with keyring tools to ensure that there is minimal exposure surface. 

Once you have your new block, it can be distributed to any distribution nodes of your choice, which will be able to verify your message is from you, and include in in their cloud. If the instruction within is merely a publish, then no further action is needed. If it is another action type, such as a head migration, it will further validate the signal within against the authors public key, and either reject or accept it depending on the result. 


## Usage

For running the webserver, you can either build and execute the docker container, which will run on port 8080 within the container. 

You can also call `go run` on the subfolder trustedtext-webserver

This will host the required endpoints, and make use of the 3 .json files in the base directory to determine the chain of blocks, and the set of peers to check against. 

To build new blocks, you can either build and sign custom blocks manually, or you can call `go run` in the subfolder trustedtext-localapp which will open a dialogue for you to add new blocks to the local chain

## Testing 

Both the core package and the webserver have tests set up. 

For the core package it's a simple `go test` whenever you wish.

For the webserver there is a script in the core directory called `integration_test_setup.sh` which will need to be run as sudo, this will set up the required docker image running in the background to test against.