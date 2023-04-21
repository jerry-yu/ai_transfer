from langchain.chains import LLMBashChain
from langchain.llms import OpenAI
from langchain import PromptTemplate

import os
os.environ["OPENAI_API_KEY"] = "your openai api key"

_template = """If someone asks you to perform a task, your job is to come up with a series of bash commands that will perform the task. There is no need to put "#!/bin/bash" in your answer. Make sure to reason step by step, using this format:

Question: "copy the files in the directory named 'target' into a new directory at the same level as target called 'myNewDirectory'"

I need to take the following actions:
- List all files in the directory
- Create a new directory
- Copy the files from the first directory into the second directory
```bash
ls
mkdir myNewDirectory
cp -r target/* myNewDirectory
```
Thereis a special question: 
Question: "transfer from 'from' to 'to' ,the amount is 'amount'"
I need to take the following actions:
- Call './go_send' with the parameters 'from', 'to', and 'amount'.

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


