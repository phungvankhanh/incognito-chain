package zkp

import (
	"errors"
	"math/big"

	"github.com/ninjadotorg/constant/privacy"
)

// PKComZeroProof contains Proof's value
type PKComZeroProof struct {
	commitmentValue *privacy.EllipticPoint //statement
	index           *byte                  //statement
	commitmentZeroS *privacy.EllipticPoint
	z               *big.Int
}

// PKComZeroWitness contains Witness's value
type PKComZeroWitness struct {
	commitmentValue *privacy.EllipticPoint //statement
	index           *byte                  //statement
	commitmentRnd   *big.Int
}

//Protocol for opening a commitment to 0 https://link.springer.com/chapter/10.1007/978-3-319-43005-8_1 (Fig. 5)

/*Protocol for opening a PedersenCommitment to 0
Prove:
	commitmentValue is PedersenCommitment value of Zero, that is statement needed to prove
	commitmentValue is calculated by Comm_ck(H,PRDNumber)
	commitmentRnd is PRDNumber, which is used to calculate commitmentValue
	s <- Zp; P is privacy.Curve base point's order, is N
	B <- Comm_ck(0,s);  Comm_ck is PedersenCommit function using public params - privacy.Curve.Params() (G0,G1...)
						but is just commit special value (in this case, special value is 0),
						which is stick with G[Index] (in this case, Index is the Index stick with commitmentValue)
						B is a.k.a commitmentZeroS
	x <- Hash(G0||G1||G2||G3||commitmentvalue) x is pseudorandom number, which could be computed easily by Verifier
	z <- rx + s; z in Zp, r is commitmentRnd
	return commitmentZeroS, z

Verify:
	commitmentValue is PedersenCommitment value of Zero, that is statement needed to prove
	commitmentValue is calculated by Comm_ck(H,PRDNumber), a.k.a A
	commitmentZeroS, z are output of Prove function, commitmentZeroS is a.k.a B
	x <- Hash(G0||G1||G2||G3||commitmentvalue)
	boolValue <- (Comm_ck(0,z) == A.x + B); in this case, A and B needed to convert to privacy.privacy.EllipticPoint
	return boolValue
)
*/
func (pro *PKComZeroProof) Init() *PKComZeroProof {
	pro.index = new(byte)
	pro.commitmentValue = new(privacy.EllipticPoint).Zero()
	pro.commitmentZeroS = new(privacy.EllipticPoint).Zero()
	pro.z = new(big.Int)
	return pro
}

func (pro *PKComZeroProof) IsNil() bool {
	if (pro.commitmentValue == nil) || (pro.commitmentZeroS == nil) || (pro.index == nil) || (pro.z == nil) {
		return true
	}
	return false
}

// Set dosomethings
func (wit *PKComZeroWitness) Set(
	commitmentValue *privacy.EllipticPoint, //statement
	index *byte, //statement
	commitmentRnd *big.Int) {
	if wit == nil {
		wit = new(PKComZeroWitness)
	}

	wit.commitmentRnd = commitmentRnd
	wit.commitmentValue = commitmentValue
	wit.index = index
}

// Bytes ...
func (pro PKComZeroProof) Bytes() []byte {
	if pro.IsNil() {
		return []byte{}
	}
	var res []byte
	res = append(pro.commitmentValue.Compress(), pro.commitmentZeroS.Compress()...)
	res = append(res, privacy.AddPaddingBigInt(pro.z, privacy.BigIntSize)...)
	res = append(res, *pro.index)
	return res
}

// SetBytes ...
func (pro *PKComZeroProof) SetBytes(bytes []byte) error {
	if pro == nil {
		pro = pro.Init()
	}

	if len(bytes) == 0 {
		return nil
	}
	if pro.commitmentValue == nil {
		pro.commitmentValue = new(privacy.EllipticPoint)
	}
	if pro.commitmentZeroS == nil {
		pro.commitmentZeroS = new(privacy.EllipticPoint)
	}
	if pro.z == nil {
		pro.z = big.NewInt(0)
	}
	if pro.index == nil {
		pro.index = new(byte)
	}

	offset := 0
	err := pro.commitmentValue.Decompress(bytes[offset : offset + privacy.CompressedPointSize])
	if err != nil {
		return privacy.NewPrivacyErr(privacy.UnexpectedErr, errors.New("Decompressed failed!"))
	}
	offset += privacy.CompressedPointSize

	err = pro.commitmentZeroS.Decompress(bytes[offset : offset + privacy.CompressedPointSize])
	if err != nil {
		return privacy.NewPrivacyErr(privacy.UnexpectedErr, errors.New("Decompressed failed!"))
	}
	offset += privacy.CompressedPointSize

	pro.z.SetBytes(bytes[offset : offset + privacy.BigIntSize])
	offset += privacy.BigIntSize

	*pro.index = bytes[offset]
	return nil
}

// Set dosomethings
func (pro *PKComZeroProof) Set(
	commitmentValue *privacy.EllipticPoint, //statement
	index *byte, //statement
	commitmentZeroS *privacy.EllipticPoint,
	z *big.Int) {

	if pro == nil {
		pro = new(PKComZeroProof)
	}
	pro.commitmentValue = commitmentValue
	pro.commitmentZeroS = commitmentZeroS
	pro.index = index
	pro.z = z
}

//Prove generate a Proof prove that the PedersenCommitment is zero
func (wit PKComZeroWitness) Prove() (*PKComZeroProof, error) {
	//var x big.Int
	//s is a random number in Zp, with p is N, which is order of base point of privacy.Curve
	sRnd := privacy.RandInt()

	//Calculate B = commitmentZeroS = comm_ck(0,s,Index)
	commitmentZeroS := privacy.PedCom.CommitAtIndex(big.NewInt(0), sRnd, *wit.index)

	//Generate challenge x in Zp
	xChallenge := GenerateChallengeFromPoint([]*privacy.EllipticPoint{wit.commitmentValue})

	//Calculate z=r*x + s (mod N)
	z := new(big.Int).Mul(wit.commitmentRnd, xChallenge)
	z.Add(z, sRnd)
	z.Mod(z, privacy.Curve.Params().N)

	proof := new(PKComZeroProof)
	proof.Set(wit.commitmentValue, wit.index, commitmentZeroS, z)
	return proof, nil
}

//Verify verify that under PedersenCommitment is zero
func (pro *PKComZeroProof) Verify() bool {
	//Generate challenge x in Zp
	xChallenge := GenerateChallengeFromPoint([]*privacy.EllipticPoint{pro.commitmentValue})

	//verifyPoint is result of A.x + B (in ECC)
	verifyPoint := pro.commitmentZeroS.Add(pro.commitmentValue.ScalarMult(xChallenge))

	//Generate Zero number
	zeroInt := big.NewInt(0)

	//Calculate comm_ck(0,z, Index)
	commitmentZeroZ := privacy.PedCom.CommitAtIndex(zeroInt, pro.z, *pro.index)

	if !commitmentZeroZ.IsEqual(verifyPoint){
		return false
	}

	return true
}
