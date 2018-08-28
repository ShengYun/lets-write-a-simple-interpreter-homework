INTEGER = "INTEGER"
PLUS = "PLUS"
MINUS = "MINUS"
MUL = "MUL"
DIV = "DIV"
LPAREN = "LPAREN"
RPAREN = "RPAREN"
EOF = "EOF"


class Token(object):
    def __init__(self, type, value):
        self.type = type
        self.value = value


class Lexer(object):
    def __init__(self, text):
        self.text = text
        self.pos = 0
        self.current_char = text[0]

    def error(self):
        raise Exception("Invalid syntax")

    def advance(self):
        self.pos += 1
        if self.pos > len(self.text)-1:
            self.current_char = None
            return Token(EOF, "EOF")
        self.current_char = self.text[self.pos]

    def skip_whitespace(self):
        while self.current_char is not None and self.current_char.isspace():
            self.advance()

    def number(self):
        res = ''
        while self.current_char is not None and self.current_char.isdigit():
            res += self.current_char
            self.advance()
        return int(res)

    def get_next_token(self):
        while self.current_char != None:

            if self.current_char.isspace():
                self.skip_whitespace()
                continue

            if self.current_char.isdigit():
                num = self.number()
                return Token(INTEGER, num)

            if self.current_char == '+':
                self.advance()
                return Token(PLUS, '+')

            if self.current_char == '-':
                self.advance()
                return Token(MINUS, '-')

            if self.current_char == '*':
                self.advance()
                return Token(MUL, '*')

            if self.current_char == '/':
                self.advance()
                return Token(DIV, '/')

            if self.current_char == '(':
                self.advance()
                return Token(LPAREN, '(')

            if self.current_char == ')':
                self.advance()
                return Token(RPAREN, ')')

            self.error()


class AstNode(object):
    pass


class Num(AstNode):
    def __init__(self, token):
        self.token = token
        self.value = token.value


class BinOp(AstNode):
    def __init__(self, left, op, right):
        self.left = left
        self.token = op
        self.op = op
        self.right = right


class UnaryOp(AstNode):
    def __init__(self, op, expr):
        self.token = op
        self.op = op
        self.expr = expr


class Parser(object):
    def __init__(self, lexer):
        self.lexer = lexer
        self.current_token = lexer.get_next_token()

    def error(self):
        raise Exception("Parser error")

    def eat(self, token_type):
        if self.current_token.type != token_type:
            self.error()
        self.current_token = self.lexer.get_next_token()

    def factor(self):
        if self.current_token.type == INTEGER:
            token = self.current_token
            self.eat(INTEGER)
            return Num(token)
        elif self.current_token.type == LPAREN:
            self.eat(LPAREN)
            node = self.expr()
            self.eat(RPAREN)
            return node
        elif self.current_token.type == PLUS:
            token = self.current_token
            self.eat(PLUS)
            return UnaryOp(token, self.factor())
        elif self.current_token.type == MINUS:
            token = self.current_token
            self.eat(MINUS)
            return UnaryOp(token, self.factor())
        else:
            self.error()


    def term(self):
        node = self.factor()
        while self.current_token is not None and self.current_token.type in (MUL, DIV):
            token = self.current_token
            if token.type == MUL:
                self.eat(MUL)
            elif token.type == DIV:
                self.eat(DIV)
            node = BinOp(node, token, self.factor())
        return node

    def expr(self):
        node = self.term()
        while self.current_token is not None and self.current_token.type in (PLUS, MINUS):
            token = self.current_token
            if token.type == PLUS:
                self.eat(PLUS)
            elif token.type == MINUS:
                self.eat(MINUS)

            node = BinOp(node, token, self.term())
        return node

    def parse(self):
        return self.expr()


class NodeVisitor(object):
    def visit(self, node):
        method_name = 'visit_' + type(node).__name__
        visitor = getattr(self, method_name, self.generic_visit)
        return visitor(node)

    def generic_visit(self, node):
        raise Exception('No visit_{} method'.format(type(node).__name__))


class Interpreter(NodeVisitor):
    def __init__(self, parser):
        self.parser = parser

    def visit_UnaryOp(self, node):
        if node.op.type == PLUS:
            return +self.visit(node.expr)
        elif node.op.type == MINUS:
            return -self.visit(node.expr)

    def visit_BinOp(self, node):
        if node.op.type == PLUS:
            return self.visit(node.left) + self.visit(node.right)
        elif node.op.type == MINUS:
            return self.visit(node.left) - self.visit(node.right)
        elif node.op.type == MUL:
            return self.visit(node.left) * self.visit(node.right)
        elif node.op.type == DIV:
            return self.visit(node.left) / self.visit(node.right)

    def visit_Num(self, node):
        return node.value

    def interpret(self):
        tree = self.parser.parse()
        return self.visit(tree)


def main():
    while True:
        try:
            text = input('calc> ')
        except EOFError:
            break
        if not text:
            continue

        lexer = Lexer(text)
        parser = Parser(lexer)
        interpreter = Interpreter(parser)
        result = interpreter.interpret()
        print(result)


if __name__ == '__main__':
    main()
