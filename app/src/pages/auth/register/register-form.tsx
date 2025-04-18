import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { useState } from "react";
import { useMutation } from "@tanstack/react-query";
import { isAxiosError } from "axios";
import { Link } from "react-router-dom";
import toast from "react-hot-toast";
import { TriangleAlert } from "lucide-react";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { FormValidationError } from "@/components/form-validation-error";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";

import { registerUser } from "@/api/register-user";

const usernameRegex = /^[a-zA-Z][a-zA-Z0-9_.]{2,19}$/;

const registerSchema = z.object({
  name: z
    .string()
    .min(2, "O nome deve ter pelo menos 2 caracteres.")
    .max(60, "O nome deve ter no máximo 60 caracteres."),
  username: z
    .string()
    .min(3, "O nome de usuário deve ter pelo menos 3 caracteres.")
    .max(20, "O nome de usuário deve ter no máximo 20 caracteres.")
    .regex(
      usernameRegex,
      "O nome de usuário deve conter apenas letras, números, underline ou ponto e começar com uma letra."
    ),
  email: z
    .string()
    .nonempty("O e-mail é obrigatório.")
    .email("O e-mail inserido não é válido."),
});

export type RegisterFormSchema = z.infer<typeof registerSchema>;

interface RegisterFormProps {
  email?: string;
}

export function RegisterForm({ email }: RegisterFormProps) {
  const [apiErrorMessage, setApiErrorMessage] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterFormSchema>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      name: "",
      username: "",
      email: email ?? "",
    },
  });

  const { mutateAsync: registerFn } = useMutation({
    mutationFn: registerUser,
  });

  async function handleRegister(data: RegisterFormSchema) {
    setApiErrorMessage(null);

    try {
      await registerFn(data);
      toast.success(
        "Conta criada! Você receberá um link no seu e-mail para acessar sua conta.",
        { duration: 6000 }
      );
    } catch (error) {
      if (isAxiosError(error)) {
        if (error.response?.status === 409) {
          const msg = error.response?.data?.message ?? "Erro de conflito.";
          setApiErrorMessage(msg);
        } else {
          toast.error("Erro ao criar conta.");
        }
      }
    }
  }

  return (
    <form
      onSubmit={handleSubmit(handleRegister)}
      className="flex flex-col gap-6"
    >
      <div className="text-center">
        <h1 className="text-2xl font-bold">Crie sua conta no Tab Notes</h1>
        <p className="text-sm text-muted-foreground mt-1">
          Já tem uma conta?{" "}
          <Link
            to="/login"
            className="underline underline-offset-4 hover:text-primary"
          >
            Entrar
          </Link>
        </p>
      </div>

      {apiErrorMessage && (
        <Alert variant="destructive">
          <TriangleAlert className="h-4 w-4" />
          <AlertTitle>Erro ao criar conta</AlertTitle>
          <AlertDescription>{apiErrorMessage}</AlertDescription>
        </Alert>
      )}

      <div className="grid gap-3">
        <Label htmlFor="name">Nome</Label>
        <Input
          id="name"
          {...register("name")}
          placeholder="Seu nome completo"
        />
        <FormValidationError message={errors.name?.message} />
      </div>

      <div className="grid gap-3">
        <Label htmlFor="username">Nome de usuário</Label>
        <Input
          id="username"
          {...register("username")}
          placeholder="ex: joao_silva"
        />
        <FormValidationError message={errors.username?.message} />
      </div>

      <div className="grid gap-3">
        <Label htmlFor="email">E-mail</Label>
        <Input
          id="email"
          type="email"
          {...register("email")}
          placeholder="seu@email.com"
        />
        <FormValidationError message={errors.email?.message} />
      </div>

      <Button type="submit" className="w-full" size="lg">
        Criar conta
      </Button>
    </form>
  );
}
