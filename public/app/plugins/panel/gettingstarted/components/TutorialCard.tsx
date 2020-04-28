import React, { FC, MouseEvent } from 'react';
import { GrafanaTheme } from '@grafana/data';
import { Icon, stylesFactory, useTheme } from '@grafana/ui';
import { css } from 'emotion';
import store from 'app/core/store';
import { cardContent, cardStyle, iconStyle } from './sharedStyles';
import { Card } from '../types';

interface Props {
  card: Card;
}

export const TutorialCard: FC<Props> = ({ card }) => {
  const theme = useTheme();
  const styles = getStyles(theme, card.done);

  return (
    <a
      className={styles.card}
      target="_blank"
      onClick={(event: MouseEvent<HTMLAnchorElement>) => handleTutorialClick(event, card)}
    >
      <div className={cardContent}>
        <div className={styles.type}>{card.type}</div>
        <div className={styles.heading}>{card.done ? 'complete' : card.heading}</div>
        <h4>{card.title}</h4>
        <div className={styles.info}>{card.info}</div>
        <Icon className={iconStyle(theme, card.done)} name={card.icon} size="xxl" />
      </div>
    </a>
  );
};

const handleTutorialClick = (event: MouseEvent<HTMLAnchorElement>, card: Card) => {
  event.preventDefault();
  const isSet = store.get(card.key);
  if (!isSet) {
    store.set(card.key, true);
  }
  window.open(card.href, '_blank');
};

const getStyles = stylesFactory((theme: GrafanaTheme, complete: boolean) => {
  const textColor = `${complete ? theme.palette.blue95 : '#FFB357'}`;
  return {
    card: css`
      ${cardStyle(theme, complete)}
      width: 460px;
      min-width: 460px;

      @media only screen and (max-width: ${theme.breakpoints.xl}) {
        min-width: 368px;
      }

      @media only screen and (max-width: ${theme.breakpoints.lg}) {
        min-width: 272px;
      }
    `,
    type: css`
      color: ${textColor};
      text-transform: uppercase;
    `,
    heading: css`
      text-transform: uppercase;
      color: ${textColor};
      margin-bottom: ${theme.spacing.sm};
    `,
    info: css`
      margin-bottom: ${theme.spacing.md};
    `,
    status: css`
      display: flex;
      justify-content: flex-end;
    `,
  };
});
