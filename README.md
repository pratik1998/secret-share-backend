# secret-share-backend

Have you ever given your password or API credentials to someone else to allow them temporary access to your accounts? If so, you're not alone. However, it's important to consider the security implications of sharing sensitive information. Ideally, passwords should only be shared with the intended person for a limited time, such as a day or a week. The main security concern is that if our communication channels (such as email, iMessage, or WhatsApp) are compromised, the attacker could gain access to other credentials.

To address this problem, I am planning to create a platform called SafeSecretShare that will allow users to securely share sensitive information such as passwords, API credentials, and authorization tokens. The platform will not store any sensitive data in plain text, but instead will use strong encryption to protect the information. The encryption key will only be shared with the owner, who can then decide who they want to share it with. This will enhance the overall privacy of the transaction.

The main features of SafeSecretShare will include
- Ability to share sensitive information securely
- Store data in encrypted format
- Allow users to set time limits for automatic deletion of data
- Limit the number of times data can be viewed
- Provide users with an access log for their sensitive data

# Requirements
`Golang`

# Installation

1. Download the git repository: 
`git clone https://github.com/pratik1998/secret-share-backend.git`

2. Download the all dependency code by running: 
`git mod tidy`

3. Start web server 
`go run .`

References:
- https://wormhole.app/ 
- https://github.com/restuwahyu13/go-rest-api
