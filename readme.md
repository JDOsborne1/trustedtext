# Trusted Text

Trusted text aims to be a distributed ecosystem for creating, sharing and updating text on the web. The intent is to be one of the tools which helps combat the ongoing loss of trust in most existing forms of information. 

Content hosted using trusted text will be modifiable, But not secretly, and the original source will always be available.

## Key objectives

* It is possible for an author to claim a segment of trusted text, by signing it with an identity key (such as PGP). 
* It is possible to identify a segment of trusted text using a set of keys, identified by the author at creation
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