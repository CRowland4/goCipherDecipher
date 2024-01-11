from hstest import *
from random import randrange


class TestUserProgram(StageTest):
    @dynamic_test
    def test(self):
        pr = TestedProgram()
        pr.start()

        g, p = randrange(100, 1000), randrange(100, 1000)

        raw_output = pr.execute(f"g is {g} and p is {p}")

        self.check_empty_or_none_output(raw_output)

        if raw_output.lower().strip() != "ok":
            raise WrongAnswer(
                "Your output should be equal to  \"OK\".\n"
                f"Your output: {raw_output}.")

        a = randrange(111690, 111690 + 100)
        A = (g ** a) % p

        raw_output = pr.execute(f"A is {A}")

        self.check_empty_or_none_output(raw_output)

        # Should get 3 lines of output back: B, A, and S
        respLines = [line.strip() for line in raw_output.split('\n') if len(line.strip()) > 0]

        if len(respLines) < 3:
            raise WrongAnswer(
                f"Expected 3 lines after receiving Alice's 'A' value, found {len(respLines)}")

        string_with_user_pub_key = respLines[0]
        string_with_comp_pub_key = respLines[1]
        string_with_shared_secret = respLines[2]

        B = 0
        try:
            B = int(string_with_user_pub_key.lower().lstrip("b is ").strip())

        except ValueError:
            raise WrongAnswer(
                "It is not possible to get B from your output.\n"
                f"Your output: {string_with_user_pub_key}.")

        try:
            int(string_with_comp_pub_key.lower().lstrip("a is ").strip())

        except ValueError:
            raise WrongAnswer(
                "It is not possible to get A from your output.\n"
                f"Your output: {string_with_comp_pub_key}.")

        userSharedSecret = None
        try:
            userSharedSecret = int(string_with_shared_secret.lower().lstrip("s is ").strip())

        except ValueError:
            raise WrongAnswer(
                "It is not possible to get S from your output.\n"
                f"Your output: {string_with_shared_secret}.")

        expectedSecret = (B ** a) % p
        if expectedSecret != userSharedSecret:
            raise WrongAnswer(
                "Your value for S does not match the expected value.\n"
                f"Given p={p}, A={A}, B={B}\n"
                f"S should be S={expectedSecret}.\n"
                f"(Note that there is no specific value for b required. Your code should work for any b you choose.)\n"
                f"This check was done versus S=(B^a) mod p , where a is the secret value selected by the test used to make A."

            )

        return CheckResult.correct()

    @staticmethod
    def check_empty_or_none_output(output):
        if not output:
            raise WrongAnswer("Your output is empty or None.")


if __name__ == '__main__':
    TestUserProgram().run_tests()
