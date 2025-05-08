export default function randomColor(): string {
  const colors = ["red", "orange", "green", "blue", "purple"];

  return colors[Math.floor(Math.random() * colors.length)];
}
