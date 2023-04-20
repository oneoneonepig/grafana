import { css } from '@emotion/css';
import React from 'react';

import { GrafanaTheme2, PageLayoutType } from '@grafana/data';
import { useStyles2, useTheme2 } from '@grafana/ui';

import { Page } from '../Page/Page';

export function PageNotFound() {
  const styles = useStyles2(getStyles);
  const theme = useTheme2();

  return (
    <Page navId="home" layout={PageLayoutType.Canvas} pageNav={{ text: 'Page not found' }}>
      <div className={styles.container}>
        <h1>Page not found</h1>
        <div className={styles.subtitle}>
          We&apos;re looking but can&apos;t seem to find this page. Try returning{' '}
          <a href="/" className="external-link">
            home
          </a>{' '}
          or seeking help on the{' '}
          <a href="https://community.grafana.com" target="_blank" rel="noreferrer" className="external-link">
            community site.
          </a>
        </div>
        <div className={styles.grot}>
          <img src={`public/img/grot-404-${theme.isDark ? 'dark' : 'light'}.svg`} width="100%" alt="grot" />
        </div>
      </div>
    </Page>
  );
}

export function getStyles(theme: GrafanaTheme2) {
  return {
    container: css({
      display: 'flex',
      flexDirection: 'column',
      padding: theme.spacing(8, 2, 2, 2),
      h1: {
        textAlign: 'center',
      },
    }),
    subtitle: css({
      color: theme.colors.text.secondary,
      fontSize: theme.typography.h5.fontSize,
      padding: theme.spacing(2, 0),
      textAlign: 'center',
    }),
    grot: css({
      maxWidth: '550px',
      paddingTop: theme.spacing(8),
      margin: '0 auto',
    }),
  };
}
