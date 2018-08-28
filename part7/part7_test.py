import unittest
import part7


class Test_TestLexer(unittest.TestCase):
    def test_lexer_simple_no_space(self):
        lexer = part7.Lexer("1+2")
        token = lexer.get_next_token()
        self.assertEqual(token.type, part7.INTEGER)
        self.assertEqual(token.value, 1)

        token = lexer.get_next_token()
        self.assertEqual(token.type, part7.PLUS)
        self.assertEqual(token.value, '+')

        token = lexer.get_next_token()
        self.assertEqual(token.type, part7.INTEGER)
        self.assertEqual(token.value, 2)

    def test_lexer_simple_with_space(self):
        lexer = part7.Lexer("   1  +  2  ")
        token = lexer.get_next_token()
        self.assertEqual(token.type, part7.INTEGER)
        self.assertEqual(token.value, 1)

        token = lexer.get_next_token()
        self.assertEqual(token.type, part7.PLUS)
        self.assertEqual(token.value, '+')

        token = lexer.get_next_token()
        self.assertEqual(token.type, part7.INTEGER)
        self.assertEqual(token.value, 2)

    def test_lexer_long_integer_with_space(self):
        lexer = part7.Lexer("   1123  /  22  ")
        token = lexer.get_next_token()
        self.assertEqual(token.type, part7.INTEGER)
        self.assertEqual(token.value, 1123)

        token = lexer.get_next_token()
        self.assertEqual(token.type, part7.DIV)
        self.assertEqual(token.value, '/')

        token = lexer.get_next_token()
        self.assertEqual(token.type, part7.INTEGER)
        self.assertEqual(token.value, 22)


if __name__ == '__main__':
    unittest.main()
