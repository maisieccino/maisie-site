import { Title } from "../typography";

export interface HeadingProps {
  title: string;
}

export const Heading = ({ title }: HeadingProps) => (
  <div className="heading">
    <Title>{title}</Title>
  </div>
)
