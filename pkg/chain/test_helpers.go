package trustedtext

const junk_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
const junk_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"

const first_standard_message = "b83030a13322e34fe61ef7dfe6d4750cab4d7429"
const second_standard_message = "f655762bf9c727eb04a71072b26e23c13b7d765c"

func helper_generate_additonal_test_block(_existing_chain Trustedtext_chain_s) Trustedtext_s {
	dexters_instruction_2 := tt_body{
		Instruction_type: "publish",
		Instruction:      "Intruder alert, DeeDee in the lab. Again!",
	}
	new_block, _ := instantiate(junk_pub_key, dexters_instruction_2, junk_pri_key)
	return new_block
}

func helper_generate_standard_test_chain(_init_only bool) Trustedtext_chain_s {
	dexters_instruction_1 := tt_body{
		Instruction_type: "publish",
		Instruction:      "Intruder alert, DeeDee in the lab",
	}
	test_ttc, _ := genesis(
		junk_pub_key,
		[]string{"lab"},
		junk_pri_key,
	)
	if _init_only {
		return test_ttc
	}

	new_block, _ := instantiate(junk_pub_key, dexters_instruction_1, junk_pri_key)

	test_ttc, _ = amend(
		test_ttc,
		new_block,
	)
	return test_ttc
}

func Test_helper_generate_standard_test_block() (Trustedtext_s, error) {
	dexters_instruction_1 := tt_body{
		Instruction_type: "publish",
		Instruction:      "DeeDee Better not interfere with this one",
	}
	return instantiate(junk_pub_key, dexters_instruction_1, junk_pri_key)
}
