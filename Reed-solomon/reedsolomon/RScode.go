package reedsolomon

import (
	"fmt"
	"number/Reed-solomon/reedsolomon/correct"
	"number/Reed-solomon/reedsolomon/field"
)

func Printf(){
	fmt.Println(field.Primitive_polynomial_8_6_4_3_2_1_0)
}

func Reed_solomon_encode(rs *correct.Correct_reed_solomon,  msg *uint8, msg_length int, encoded *uint8) int {
if (msg_length > rs.message_length) {
return -1;
}

size_t pad_length = rs->message_length - msg_length;
for (unsigned int i = 0; i < msg_length; i++) {
// message goes from high order to low order but libcorrect polynomials go low to high
// so we reverse on the way in and on the way out
// we'd have to do a copy anyway so this reversal should be free
rs->encoded_polynomial.coeff[rs->encoded_polynomial.order - (i + pad_length)] = msg[i];
}

// 0-fill the rest of the coefficients -- this length will always be > 0
// because the order of this poly is block_length and the msg_length <= message_length
// e.g. 255 and 223
memset(rs->encoded_polynomial.coeff + (rs->encoded_polynomial.order + 1 - pad_length), 0, pad_length);
memset(rs->encoded_polynomial.coeff, 0, (rs->encoded_polynomial.order + 1 - rs->message_length));

polynomial_mod(rs->field, rs->encoded_polynomial, rs->generator, rs->encoded_remainder);

// now return byte order to highest order to lowest order
for (unsigned int i = 0; i < msg_length; i++) {
encoded[i] = rs->encoded_polynomial.coeff[rs->encoded_polynomial.order - (i + pad_length)];
}

for (unsigned int i = 0; i < rs->min_distance; i++) {
encoded[msg_length + i] = rs->encoded_remainder.coeff[rs->min_distance - (i + 1)];
}

return rs->block_length;
}
