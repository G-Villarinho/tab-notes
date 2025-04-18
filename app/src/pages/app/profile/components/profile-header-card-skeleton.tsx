import { Skeleton } from "@/components/ui/skeleton";

export function ProfileHeaderCardSkeleton() {
  return (
    <div className="rounded-xl border shadow-sm bg-card overflow-hidden">
      <div className="h-36 w-full bg-muted relative" />

      <div className="flex flex-col px-6 pt-4 pb-6 -mt-14">
        <div className="flex justify-between items-center relative">
          <Skeleton className="h-24 w-24 rounded-full border-4 border-background" />

          <Skeleton className="h-9 w-28 rounded-md absolute top-4 right-4 mt-10" />
        </div>

        <div className="flex flex-col justify-center mt-3">
          <div className="flex items-center gap-2">
            <Skeleton className="h-6 w-40 rounded-md" />
            <Skeleton className="h-4 w-4 rounded-full" />
          </div>

          <Skeleton className="h-4 w-28 mt-2" />

          <div className="flex gap-6 pt-3">
            <Skeleton className="h-3 w-20" />
            <Skeleton className="h-3 w-20" />
          </div>
        </div>
      </div>
    </div>
  );
}
