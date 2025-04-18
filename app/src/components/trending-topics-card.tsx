export function TrendingTopicsCard() {
  const trends = [
    { topic: "#SaúdeMental", posts: "3.4k posts" },
    { topic: "#Produtividade", posts: "2.6k posts" },
    { topic: "#Livros", posts: "2.1k posts" },
    { topic: "#TrabalhoRemoto", posts: "1.7k posts" },
    { topic: "#BemEstar", posts: "1.5k posts" },
    { topic: "#Rotina", posts: "1.2k posts" },
    { topic: "#Comunidade", posts: "1.1k posts" },
  ];

  return (
    <div className="rounded-xl border p-4 shadow-sm bg-card space-y-4">
      <h2 className="text-sm font-semibold text-muted-foreground">
        Tendências para você
      </h2>

      <ul className="space-y-3">
        {trends.map((trend, i) => (
          <li
            key={i}
            className="flex flex-col gap-0.5 text-sm cursor-pointer hover:bg-muted p-2 rounded-md transition"
          >
            <span className="font-medium text-foreground">{trend.topic}</span>
            <span className="text-xs text-muted-foreground">{trend.posts}</span>
          </li>
        ))}
      </ul>
    </div>
  );
}
