import { Link, useNavigate } from "react-router-dom";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import toast from "react-hot-toast";
import { isAxiosError } from "axios";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { FormValidationError } from "@/components/form-validation-error";
import { useMutation } from "@tanstack/react-query";
import { login } from "@/api/login";

const loginSchema = z.object({
  email: z
    .string()
    .nonempty("O e-mail é obrigatório.")
    .email("O e-mail inserido não é válido."),
});

export type LoginFormSchema = z.infer<typeof loginSchema>;

export function LoginForm() {
  const navigate = useNavigate();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormSchema>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
    },
  });

  const { mutateAsync: loginFn } = useMutation({
    mutationFn: login,
  });

  async function handleLogin(data: LoginFormSchema) {
    try {
      await loginFn(data);
      toast.success(
        "Um link de acesso foi enviado para o seu e-mail. Verifique sua caixa de entrada.",
        {
          duration: 5000,
        }
      );
    } catch (error) {
      if (isAxiosError(error)) {
        if (error.status === 404) {
          navigate("/login/account-not-found", {
            state: { email: data.email },
          });
        } else {
          toast.error("Ocorreu um erro ao tentar fazer login.");
        }
      }
    }
  }

  return (
    <div className="flex flex-col gap-6">
      <div className="text-center">
        <h1 className="text-2xl font-bold">Entrar no Tab Notes</h1>
        <p className="text-sm text-muted-foreground mt-1">
          Ainda não tem uma conta?{" "}
          <Link
            to="/register"
            className="underline underline-offset-4 hover:text-primary"
          >
            Cadastre-se
          </Link>
        </p>
      </div>

      <form
        onSubmit={handleSubmit(handleLogin)}
        className="flex flex-col gap-6"
      >
        <div className="grid gap-3">
          <Label htmlFor="email">E-mail</Label>
          <div>
            <Input
              id="email"
              type="email"
              placeholder="fulano@email.com"
              {...register("email")}
            />
            <FormValidationError message={errors.email?.message} />
          </div>
        </div>

        <Button type="submit" className="w-full" size="lg">
          Entrar
        </Button>
      </form>

      <p className="text-muted-foreground *:[a]:hover:text-primary text-center text-xs text-balance *:[a]:underline *:[a]:underline-offset-4">
        Ao continuar, você concorda com os nossos{" "}
        <Link to="/termos">Termos de Uso</Link> e{" "}
        <Link to="/privacidade">Política de Privacidade</Link>.
      </p>
    </div>
  );
}
