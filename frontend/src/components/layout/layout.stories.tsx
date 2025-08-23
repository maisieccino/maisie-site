import type { Meta, StoryObj, Parameters } from "@storybook/react";

import { Container, Heading, BentoContainer, Bento } from ".";
import { P } from "../typography";

const meta: Meta<typeof Container> = {
  component: Container,
};

export default meta;

type Story = StoryObj<typeof Container>;

export const ContainerComponent: Story = {
  render: () => (
    <Container>
      <Heading title="test title"></Heading>
      <P>Hello I am some test text, I love coffee and more coffee. This is a long paragraph that should wrap onto a new line.</P>
      <P>Another paragraph</P>
      <BentoContainer>
        <Bento>
          Hello
        </Bento>
        <Bento title="Projects" wide>
          A recent project...
        </Bento>
        <Bento title="Data" wide color="#fd98ce">
          Maisie's data sets.
        </Bento>
      </BentoContainer>
    </Container>
  ),
};
