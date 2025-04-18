import Icon from "@/assets/icon.svg";
import { cn } from "@/lib/utils";

export function AppLoadingScreen() {
  return (
    <div className="flex items-center justify-center h-screen w-full bg-background">
      <img
        src={Icon}
        alt="Tab Notes"
        className={cn(
          "w-12 h-12 animate-pulse",
          "transition-opacity duration-300 ease-in-out"
        )}
      />
    </div>
  );
}
