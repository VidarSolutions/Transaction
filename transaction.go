package Transaction
import (
	"crypto/ed25519"

)
type Transaction struct{
	TID			uint64							//Transaction Id  first 4 digits are the year
	From		uint64							//BundleID as PublicKey where funds are coming from
	To			uint64							//BundleID as PublicKey where funds are going to
	Sender		ed25529.PublicKey				//Sender and signer of transaction public key that is a one time use just for this trasnaction
	Password	string							//Password of current owner.  Publicky an dpassword are hashed and compared to the signed hash	
	Value		uint64							//Value of this transaction
	LTID		uint64							//Transaction Id from the last transaction sent from this bundle
	Phash		bytes32							//Hash from last transaction
	Hash		bytes32							//Hash for next password and new publickey for sending bundle that prove ownership of bundle
	Sig			bytes64							//Signature of new keys that own the bundle sending the trasaction
	SID			uint64							//Number of transactions sent by this sender not including this transaction  
	RID			uint64							//Number of transactions receieved by this receiever not including this transaction  
	Approved	bool							//If transaction has been approved
	Approver	ed25529.PublicKey				//PublicKey of Approver  Must be from an authorized approver for this bundle
	Signed		bytes64							//Hash of transaction signed by approver

}


var(
	
)

func getTransaction(txID uint64)(Transaction)	{
	return Transactions[txID]

}