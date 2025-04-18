import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";

import { useAuth } from "@/hooks/use-auth";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { cn } from "@/lib/utils";
import { useMutation } from "@tanstack/react-query";
import { updateProfile } from "@/api/update-profile";
import { Save, AlertCircle } from "lucide-react";
import { UserAvatar } from "@/components/user-avatar";
import { useState } from "react";
import { isAxiosError } from "axios";

const updateProfileSchema = z.object({
  name: z.string().min(1, "O nome é obrigatório."),
  username: z.string().min(1, "O nome de usuário é obrigatório."),
});

type UpdateProfileSchema = z.infer<typeof updateProfileSchema>;

interface UpdateProfileProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

export function UpdateProfileModal({ open, onOpenChange }: UpdateProfileProps) {
  const { user, setUser } = useAuth();
  const [apiError, setApiError] = useState<string | null>(null);

  if (!user) {
    throw new Error("User not found");
  }

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<UpdateProfileSchema>({
    resolver: zodResolver(updateProfileSchema),
    defaultValues: {
      name: user.name,
      username: user.username,
    },
  });

  const { mutateAsync: updateProfileFn } = useMutation({
    mutationFn: updateProfile,
  });

  async function handleUpdateProfile(data: UpdateProfileSchema) {
    if (!user) {
      return;
    }
    setApiError(null);
    try {
      await updateProfileFn(data);
      user.name = data.name;
      user.username = data.username;
      setUser(user);
      onOpenChange(false);
    } catch (error) {
      if (isAxiosError(error) && error.response?.status === 409) {
        setApiError("Este nome de usuário já está em uso.");
      } else {
        setApiError("Erro ao atualizar perfil. Tente novamente.");
      }
    }
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[500px] rounded-xl bg-card p-0 overflow-hidden">
        <div className="h-16 w-full bg-muted" />
        <div className="-mt-8 flex justify-center">
          <UserAvatar
            name={user.name}
            username={user.username}
            className="h-16 w-16 border-4 border-card"
          />
        </div>

        <div className="px-6 pt-4 pb-6 space-y-6">
          <DialogHeader className="text-center">
            <DialogTitle className="text-xl font-semibold">
              Editar perfil
            </DialogTitle>
            <DialogDescription className="text-muted-foreground text-sm">
              Atualize suas informações públicas
            </DialogDescription>
          </DialogHeader>

          {/* Alerta de erro */}
          {apiError && (
            <Alert variant="destructive">
              <AlertCircle className="h-4 w-4" />
              <AlertTitle>Falha ao atualizar perfil</AlertTitle>
              <AlertDescription>{apiError}</AlertDescription>
            </Alert>
          )}

          <form
            onSubmit={handleSubmit(handleUpdateProfile)}
            className="grid gap-5"
          >
            <div className="space-y-2">
              <Label htmlFor="name">Nome</Label>
              <Input
                id="name"
                placeholder="Seu nome"
                {...register("name")}
                className={cn(
                  "transition-all",
                  errors.name &&
                    "border-destructive focus-visible:ring-destructive"
                )}
              />
              {errors.name && (
                <p className="text-xs text-destructive">
                  {errors.name.message}
                </p>
              )}
            </div>

            <div className="space-y-2">
              <Label htmlFor="username">Username</Label>
              <Input
                id="username"
                placeholder="seu_usuario"
                {...register("username")}
                className={cn(
                  "transition-all",
                  errors.username &&
                    "border-destructive focus-visible:ring-destructive"
                )}
              />
              {errors.username && (
                <p className="text-xs text-destructive">
                  {errors.username.message}
                </p>
              )}
            </div>

            <DialogFooter className="pt-2">
              <Button
                type="submit"
                className="w-full text-sm font-medium gap-2"
                disabled={isSubmitting}
              >
                <Save className="w-4 h-4" />
                {isSubmitting ? "Salvando..." : "Salvar alterações"}
              </Button>
            </DialogFooter>
          </form>
        </div>
      </DialogContent>
    </Dialog>
  );
}
