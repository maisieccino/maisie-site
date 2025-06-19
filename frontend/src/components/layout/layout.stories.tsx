import type { Meta, StoryObj, Parameters } from "@storybook/react";

import { Container } from ".";
import { Title } from "../typography";

const meta: Meta<typeof Container> = {
  component: Container,
};

export default meta;

type Story = StoryObj<typeof Container>;

export const ContainerComponent: Story = {
  render: () => (
    <Container>
      <Title>Test Title</Title>
      <p>Hello I am some test text, I love coffee and more coffee</p>
    </Container>
  ),
};
