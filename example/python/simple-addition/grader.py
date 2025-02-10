import json
import sys

from student import add, subtract, divide, multiply

def grade_student_code(student_code_filename):
    """Grades student code by importing and testing functions."""

    results = {
        "addition": {"passed": False, "message": ""},
        "subtraction": {"passed": False, "message": ""},
        "multiplication": {"passed": False, "message": ""},
        "division": {"passed": False, "message": ""},
    }

    try:
        # Dynamically import the student's code

        # Test addition
        try:
            student_addition = add  # Assuming function name is 'add'
            result = student_addition(5, 3)
            if result == 8:
                results["addition"]["passed"] = True
            else:
                results["addition"]["message"] = f"Addition failed. Expected 8, got {result}"
        except Exception as e:
            results["addition"]["message"] = f"Addition test caused an error: {e}"

        # Test subtraction
        try:
            student_subtraction = subtract
            result = student_subtraction(10, 4)
            if result == 6:
                results["subtraction"]["passed"] = True
            else:
                results["subtraction"]["message"] = f"Subtraction failed. Expected 6, got {result}"
        except Exception as e:
            results["subtraction"]["message"] = f"Subtraction test caused an error: {e}"

        # Test multiplication
        try:
            student_multiplication = multiply
            result = student_multiplication(6, 7)
            if result == 42:
                results["multiplication"]["passed"] = True
            else:
                results["multiplication"]["message"] = f"Multiplication failed. Expected 42, got {result}"
        except Exception as e:
            results["multiplication"]["message"] = f"Multiplication test caused an error: {e}"

        # Test division
        try:
            student_division = divide
            result = student_division(20, 5)
            if result == 4:
                results["division"]["passed"] = True
            else:
                results["division"]["message"] = f"Division failed. Expected 4, got {result}"
        except Exception as e:
            results["division"]["message"] = f"Division test caused an error: {e}"

    except Exception as e:
        return json.dumps({"error": f"Error importing or executing student code: {e}"})

    return json.dumps(results)


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python grader.py <student_code_filename>")
        sys.exit(1)

    student_code_file = sys.argv[1]
    final = grade_student_code(student_code_file)
    print(final)
