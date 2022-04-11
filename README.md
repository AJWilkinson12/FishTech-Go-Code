# FishTech-Go-Code
Coding challenge for FishTech interview. Using Golang, I created a simple API that is deployed as a serverless function on a cloud platform using Terraform.
This API is a IP Information collector. Using it will give you various information about a given IP. Just insert it at the end of the link given below.
I've included the already zipped file for convenience if someone who like to Terraform Apply using that zip file and creating another link.

## Notes for reviewers.
I had initial problems getting Terraform to get the function uploaded and a couple times it told me there was an interal server error. But after doing research
and watching multiple videos on deployment using Terraform I was able to come to a solution. Terraform is an incredible tool and I had fun learning about it
and using it to deploy this serverless funtion. I am excited to have more reasons to keep working with it! The following link will take you to the API
and the IP inserted already is an University of Missouri IP Address: https://szjio5wb9k.execute-api.us-east-1.amazonaws.com/v1/fishTech?ip=161.130.189.199

Here is the link again without the inserted IP: `https://szjio5wb9k.execute-api.us-east-1.amazonaws.com/v1/fishTech?ip={INSERT DESIRED IP}`
