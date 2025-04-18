import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";
import toast from "react-hot-toast";
import { isAxiosError } from "axios";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { FormValidationError } from "@/components/form-validation-error";
import { useMutation } from "@tanstack/react-query";
import { login } from "@/api/login";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";

const loginSchema = z.object({
  email: z
    .string()
    .nonempty("O e-mail é obrigatório.")
    .email("O e-mail inserido não é válido."),
});

export type LoginFormSchema = z.infer<typeof loginSchema>;

interface LoginAccountNotFoundFormProps {
  latestEmail: string;
}

export function LoginAccountNotFoundForm({
  latestEmail,
}: LoginAccountNotFoundFormProps) {
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

  async function handleRetry(data: LoginFormSchema) {
    try {
      await loginFn(data);
      toast.success(
        "Um link de acesso foi enviado para o seu e-mail. Verifique sua caixa de entrada."
      );
      navigate("/login");
    } catch (error) {
      if (isAxiosError(error) && error.response?.status === 404) {
        navigate("/login/account-not-found", {
          state: { email: data.email },
        });
      } else {
        toast.error("Erro ao tentar reenviar o link.");
      }
    }
  }

  return (
    <div className="flex flex-col gap-6">
      <div className="text-center">
        <h1 className="text-2xl font-bold">Conta não encontrada</h1>
        <p className="text-sm text-muted-foreground mt-1">
          Não encontramos nenhuma conta com o e-mail{" "}
          <strong>{latestEmail}</strong>.
        </p>
      </div>

      <form
        onSubmit={handleSubmit(handleRetry)}
        className="flex flex-col gap-6"
      >
        <div className="grid gap-3">
          <Label htmlFor="email">Tentar com outro e-mail</Label>
          <Input
            id="email"
            type="email"
            placeholder="seuemail@email.com"
            {...register("email")}
          />
          <FormValidationError message={errors.email?.message} />
        </div>

        <div className="flex gap-2">
          <Button type="submit" className="flex-1" size="lg">
            Tentar novamente
          </Button>
          <Button
            type="button"
            variant="outline"
            className="flex-1"
            size="lg"
            onClick={() =>
              navigate("/register", { state: { email: latestEmail } })
            }
          >
            Criar conta
          </Button>
        </div>
      </form>

      <p className="text-muted-foreground *:[a]:hover:text-primary text-center text-xs text-balance *:[a]:underline *:[a]:underline-offset-4">
        Ao continuar, você concorda com os nossos{" "}
        <Link to="/termos">Termos de Uso</Link> e{" "}
        <Link to="/privacidade">Política de Privacidade</Link>.
      </p>
    </div>
  );
}
