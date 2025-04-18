import { AlertTriangle } from "lucide-react";

interface FormValidationErrorProps {
  message?: string;
}

export function FormValidationError({ message }: FormValidationErrorProps) {
  if (!message) {
    return null;
  }

  return (
    <div className="flex items-center gap-1 text-md text-red-500">
      <AlertTriangle className="size-4" />
      <small className="mt-1">{message}</small>
    </div>
  );
}
