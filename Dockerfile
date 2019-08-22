FROM python
RUN pip install tinydb
RUN pip install telepot
ADD . /
CMD ["python","-u", "/main.py"]