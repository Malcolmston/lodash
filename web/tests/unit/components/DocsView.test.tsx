import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import { DocsView } from '../../../src/components/DocsView';
import type { DocIndex } from 'go-ui';

// A minimal DocIndex the stubbed fetch returns for DocsApp's doc.json request.
const DOC_INDEX: DocIndex = {
  module: 'github.com/malcolmston/lodash',
  packages: [
    {
      importPath: 'github.com/malcolmston/lodash',
      name: 'lodash',
      synopsis: 'Package lodash is a standard-library-only functional utility library for Go.',
      doc: 'Package lodash is a standard-library-only functional utility library for Go.',
      consts: [],
      vars: [],
      types: [
        {
          name: 'Debouncer',
          signature: 'type Debouncer struct{}',
          doc: 'Debouncer coalesces rapid calls into a single delayed invocation.',
          consts: [],
          vars: [],
          funcs: [],
          methods: [],
        },
      ],
      funcs: [{ name: 'Filter', signature: 'func Filter[T any](s []T, fn func(T) bool) []T', doc: 'Filter keeps elements for which fn returns true.' }],
    },
  ],
};

describe('DocsView', () => {
  beforeEach(() => {
    // DocsApp fetches doc.json; return the small index.
    global.fetch = vi.fn((input: RequestInfo | URL) => {
      if (String(input).includes('doc.json')) {
        return Promise.resolve({ ok: true, json: () => Promise.resolve(DOC_INDEX) } as Response);
      }
      return new Promise<Response>(() => {});
    }) as unknown as typeof fetch;
  });

  it('renders the inline React API reference from the fetched doc.json', async () => {
    const { container } = render(<DocsView />);
    expect(container.querySelector('#view-docs')).not.toBeNull();
    expect(
      screen.getByRole('heading', { level: 2, name: /API documentation/ }),
    ).toBeInTheDocument();

    // DocsApp fetches asynchronously, then renders the package view + symbols.
    expect(await screen.findByRole('heading', { name: /package lodash/ })).toBeInTheDocument();
    expect(container.querySelector('#sym-Filter'), 'func Filter symbol card').not.toBeNull();
    expect(container.querySelector('#sym-Debouncer'), 'type Debouncer symbol card').not.toBeNull();

    // The secondary link to the raw generated static HTML remains.
    expect(screen.getByRole('link', { name: /Open the raw generated HTML/ })).toHaveAttribute('href', './api/');
  });
});
