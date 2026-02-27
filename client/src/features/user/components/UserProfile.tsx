export const UserProfile = () => {
  const user = {
    name: "Erzhan",
    role: "Администратор",
  };
  return (
    <div className="flex items-center gap-2">
      <div className="w-10 h-10 bg-gray-400 rounded-xl"></div>
      <div className="flex flex-col">
        <span className="font-semibold">{user.name}</span>
        <span className="text-sm">{user.role}</span>
      </div>
    </div>
  );
};
