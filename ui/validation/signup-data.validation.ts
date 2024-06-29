import { z } from 'zod';

// Define the password regex for the required criteria
const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}$/;

// Define the username regex for the required criteria
const usernameRegex = /^[a-zA-Z0-9_.]+$/;

// Define the fullName regex for the required criteria
const fullNameRegex = /^[A-Za-z ]+$/;

// Zod schema for the form validation
const signUpFormValidation = z.object({
  fullName: z.string()
    .trim()
    .regex(fullNameRegex, { message: "Full name must contain only alphabets and spaces" })
    .min(1, { message: "Full name cannot be empty" })
    .refine((val) => !val.startsWith(' ') && !val.endsWith(' '), { message: "Full name cannot start or end with a space" }),
  username: z.string()
    .regex(usernameRegex, { message: "Username can only contain letters, numbers, underscores, and dots" })
    .min(1, { message: "Username cannot be empty" }),
  email: z.string()
    .email({ message: "Invalid email address" }),
  password: z.string()
    .regex(passwordRegex, { message: "Password must be at least 8 characters long and contain uppercase, lowercase, number, and special character" }),
  confirmPassword: z.string()
}).refine(data => data.password === data.confirmPassword, {
  message: "Passwords do not match",
  path: ["confirmPassword"]
});
export default signUpFormValidation;
export type signUpFormValidationType = z.infer<typeof signUpFormValidation>