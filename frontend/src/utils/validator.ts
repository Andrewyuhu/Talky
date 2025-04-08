function isNotBlank(s: string) {
  return s.length > 0;
}

function isValidEmail(s: string) {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(s);
}

function addError(
  valid: boolean,
  errorField: string,
  errorMessage: string,
  errors: Map<string, string>
) {
  if (valid || errors.has(errorField)) {
    return;
  }
  errors.set(errorField, errorMessage);
}

function isMinLength(s: string, min: number): boolean {
  return s.length > min;
}

function isMaxLength(s: string, max: number): boolean {
  return s.length < max;
}

function isAlphanumeric(s: string): boolean {
  const alphanumericRegex = /^[a-zA-Z0-9]+$/;
  return alphanumericRegex.test(s);
}

export {
  isNotBlank,
  isValidEmail,
  addError,
  isMinLength,
  isMaxLength,
  isAlphanumeric,
};
