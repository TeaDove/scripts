import streamlit as st

"""
# Header
## Header 2
### Header 3
Regular text with `code`, *bald* and **italic**!
``` bash
# Bash code!
ls -la
ping localhost
```
``` python
# Python3 code!
def main():
    a = 2
    print(a ** a)
```
"""

if st.sidebar.checkbox("Show cpp code"):
    """
    ``` c
    int main()
        // some code in c
        int a = 2;
        int *b = a;
        return &b + a;
    ```
    """
if st.sidebar.button("Say hello"):
    """
    Hello!
    """
else:
    """
    Bye
    """
