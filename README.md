# FaceMatch AI

## How it works

First step: you provide your Instagram account. From here, all your 1st degree connections will be scanned using the LinkedIn API. You will be required to sign in using LinkedIn so that the APU has access to your profile.

Then, your connections' profile pictures and yours will be serailized, creating embeddings which are used for the face matching. The image is standardized to a 1x1 480px format and converted to base64. This value is sent to AWS Bedrock using the Titan model.

This embedding and basic profile data is added to the database.
