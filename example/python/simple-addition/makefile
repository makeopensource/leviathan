PYTHON = python3  # Or python3, depending on your system
GRADER_SCRIPT = grader.py  # Name of your grader script

# Target to grade the student's code.
grade:
	$(PYTHON) $(GRADER_SCRIPT) student.py

all:
	make incorrect
	make correct
	make timeout

correct:
	$(PYTHON) $(GRADER_SCRIPT) student_correct.py

incorrect:
	$(PYTHON) $(GRADER_SCRIPT) student_incorrect.py


timeout:
	$(PYTHON) $(GRADER_SCRIPT) student_timeout.py
