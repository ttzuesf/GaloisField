package correct

import (
	"github.com/number/Reed-solomon/reedsolomon/field"
)

type Correct_reed_solomon struct {
	Block_length int;
	Message_length int;
	Min_distance int;

	first_consecutive_root field.Field_logarithm_t ;
	generator_root_gap field.Field_logarithm_t ;

	field field.Field_t ;

	generator field.Polynomial_t ;
	generator_roots *field.Field_element_t ;
	generator_root_exp **field.Field_logarithm_t;

	encoded_polynomial field.Polynomial_t;
	encoded_remainder field.Polynomial_t;

	syndromes *field.Field_element_t;
	modified_syndromes *field.Field_element_t;
	received_polynomial field.Polynomial_t ;
	error_locator field.Polynomial_t ;
	error_locator_log field.Polynomial_t;
	erasure_locator field.Polynomial_t ;
	error_roots *field.Field_element_t;
	error_vals *field.Field_element_t;
	error_locations *field.Field_logarithm_t;

	element_exp **field.Field_logarithm_t;
	// used during find_error_locator
	last_error_locator field.Polynomial_t;

	// used during error value search
	error_evaluator field.Polynomial_t;
	error_locator_derivative field.Polynomial_t;
	init_from_roots_scratch[2] field.Polynomial_t;
	has_init_decode bool;
};


func(rs *Correct_reed_solomon) Reed_solomon_encode(msg []uint8, encoded []uint8) int {
	msg_length:=len(msg);
	if (msg_length > rs.Message_length) {
		return -1;
	}
	rs.encoded_polynomial.Coeff=make([]field.Field_element_t,rs.Message_length);
	var pad_length int= rs.Message_length - msg_length;
	for i:= 0; i < msg_length; i++ {
		// message goes from high order to low order but libcorrect polynomials go low to high
		// so we reverse on the way in and on the way out
		// we'd have to do a copy anyway so this reversal should be free
		rs.encoded_polynomial.Coeff[rs.encoded_polynomial.Order - (i + pad_length)] = field.Field_element_t(msg[i]);
	}

	// 0-fill the rest of the coefficients -- this length will always be > 0
	// because the order of this poly is block_length and the msg_length <= message_length
	// e.g. 255 and 223
	encoded=make([]uint8,rs.Block_length);
	copy(encoded[0:len(msg)],msg)

	polynomial_mod(rs.field, rs.encoded_polynomial, rs.generator, rs.encoded_remainder);

	// now return byte order to highest order to lowest order
	for i:= 0; i < msg_length; i++ {
		encoded[i] = rs.encoded_polynomial.Coeff[rs.encoded_polynomial.Order - (i + pad_length)];
	}

	for (unsigned int i = 0; i < rs->min_distance; i++) {
		encoded[msg_length + i] = rs->encoded_remainder.coeff[rs->min_distance - (i + 1)];
	}

	return rs->block_length;
}