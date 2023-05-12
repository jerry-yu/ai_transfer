from langchain.chains import LLMBashChain
from langchain.llms import OpenAI
from langchain import PromptTemplate

_template = """Prefix: dpr=100000000000000000,all your answer start with "```bash",'send' is sub-command of program; If someone asks you to perform a task, you need to call the X program accordingly, with different parameters. No need to put "#!/bin/bash" in your answer. Make sure to reason step-by-step using the following format:

Question: "send 100 dpr from 'from' to 'to'"

- if 'from' is one of [Alice, Bob, Eve], take the following actions:
```bash
./X send 'from' 'to' 100000000000000000000
```

- if 'from' is not one of [Alice, Bob, Eve], take the following actions:
```bash
./X send 'from' 'to' 100000000000000000000 'from'.json
```

- there is special question which does not involve transferring or sending something,take the following actions,do not add any symbol in your answer:
```bash
You can only tranfer tokens here!
```

That is the format. Begin!

Question: {question}"""

llm = OpenAI(temperature=0)

#text = "send from Alice to Bob ,the amount is 99990000000000000000000"

multiple_input_prompt = PromptTemplate(
    input_variables=["question"], 
    template= _template
)
#multiple_input_prompt.format(adjective="Alice", content="Bob")

bash_chain = LLMBashChain(llm=llm, prompt=multiple_input_prompt, verbose=True)

text = input("input your request\n")
try:
    while True:
        print("your input is: ", text)
        bash_chain.run(text)
        text = input("input your request\n")

except KeyboardInterrupt:
    print("exit")


