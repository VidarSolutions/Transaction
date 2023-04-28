package Transaction
import (

	"crypto/ed25519"

)

type Bytes32 [32]byte
type Bytes64 [64]byte


type Tx struct{
	TID			uint64							//Transaction Id  first 4 digits are the year
	From		uint64							//BundleID as PublicKey where funds are coming from
	To			uint64							//BundleID as PublicKey where funds are going to
	Sender		ed25529.PublicKey				//Sender and signer of transaction public key that is a one time use just for this trasnaction
	Password	string							//Password of current owner.  Publicky an dpassword are hashed and compared to the signed hash	
	Value		uint64							//Value of this transaction
	LTID		uint64							//Transaction Id from the last transaction sent from this bundle
	Phash		appData.bytes32					//Hash from last transaction
	Hash		appData.bytes32					//Hash for next password and new publickey for sending bundle that prove ownership of bundle
	Sig			appData.bytes64					//Signature of new keys that own the bundle sending the trasaction
	SID			uint64							//Number of transactions sent by this sender not including this transaction  
	RID			uint64							//Number of transactions receieved by this receiever not including this transaction  
	Approved	bool							//If transaction has been approved
	Approver	ed25529.PublicKey				//PublicKey of Approver  Must be from an authorized approver for this bundle
	Signed		appData.bytes64					//Hash of transaction signed by approver


}

var(

//setup map of transactions
	Transactions		map[uint64]Transaction			//Known Transactions
	notValid  			map[uint64][]transaction		//Map of potential transactions ID to node IDs indicating a bad transaction was approved. We do not track valid transaction in a map that is track in each bundle, but this is a way to track and vote a bad transaction be reversed and removed.
	authorized  		map[uint64][]uint64				//Map of bundle IDs to a list of Authorized Node IDs

)


func getTransaction(txID uint64)(Tx)	{
	return Transactions[txID]

}